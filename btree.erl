-module(btree).
%-export([open/1, insert/2, lookup/1]).
-export([open/1]).

-define(KP_NODE, 1).
-define(KV_NODE, 2).
-define(MAX_NODE, 4).
-define(LOG, io:format).
-define(FILE_PATH, "/Users/nimishgupta/test1").
-record(node, {type, key, value}).
-record(accnode, {count, childoffset}).
-record(kvnodevalue, {value, start_key, end_key}).

generatedata(Num) ->
    generatedata([], Num).

generatedata(L, 0) ->
    L;
generatedata(L, Num) ->
    Key = "key" ++ integer_to_list(Num),
    Value = "value" ++ integer_to_list(Num),
    L2 = L ++ [{Key,Value}],
    generatedata(L2, Num-1). 

getlen(Ll) ->
    lists:foldl(fun(X, Count) -> Count = Count+1 end, 0, Ll).

write_node_to_disk(Key, Value, Type, Fd) ->
    NewNode = #node{type=Type,key=Key,value=Value},
    Data = base64:encode_to_string(term_to_binary(NewNode)),
    {ok, Offset} = file:position(Fd, eof),
    case file:write(Fd, Data) of 
       ok -> 
           ?LOG("wrote to the disk",[]),
           Offset;
       {error, _} ->
           ?LOG("Failed to write to disk", []),
           -1
    end.

build_node(Data, 0, MyList, _)->
    MyList;
build_node(Data, ElementCount, MyList, Fd) ->
    case lists:nth(ElementCount, Data) of
        {Key, Value} ->
            Offset = write_node_to_disk(Key, Value, ?KP_NODE, Fd);
        Value ->
            Offset = write_node_to_disk("", Value, ?KV_NODE, Fd)
        end,
    if 
        Offset >= 0 ->
            MyList1 = MyList ++ [Offset],         
            ?LOG("successfully written to disk offset ~p", [Offset]),
            build_node(Data, ElementCount-1, MyList1, Fd);
        true ->
            ?LOG("failed to write to disk", [])
    end.

build_new_list(MyList, ResultList, Count) ->
    ?LOG("build new list called", []),
    Value = lists:sublist(MyList, Count+1, ?MAX_NODE),
    ResultList1 = ResultList ++ [Value],
    Count1 = Count+?MAX_NODE,
    Len = length(MyList),
    if 
        Count1 < Len ->
            build_new_list(MyList, ResultList1, Count1);
        Count1 >= Len ->
            ResultList1
    end.

%build_new_value_list(MyList, ResultList) ->
 %   lists:foldl(fun(X, Acc) ->
  %              if 
   %                 Acc%?MAX_NODE = 0 -> 
    %                    Value = lists:sublist(MyList, Acc, ?MAX_NODE),
    %                    ResultList1 = ResultList ++ [Value]; 
    %        end,
    %        V = Acc#accnode.count,
    %        K = V+1,

build_recursive(Data, Fd, MyList) ->
    Result = build_node(Data, length(Data), MyList, Fd),
    case length(Result) of 
        1 ->
            ?LOG("final result is ~p", [Result]),
            Result;
        _ ->
             ?LOG("list is ~p", [Result]),
             NewList = build_new_list(Result, [], 0),
             ?LOG("newlist is ~p", [NewList]),
             build_recursive(NewList, Fd, [])   
    end.

build_recursive(Data, Fd) ->
    build_recursive(Data, Fd, []).

read_tree(Fd, Offset) ->
    V = file:pread(Fd, 36, 36),
    case V of
        {ok, Data} ->
            Read = binary_to_term(base64:decode(Data)),
            ?LOG("~p",[Read#node.key]);
        eof ->
            ?LOG("Error in reading from file",[])
    end.

build_tree(Data, Fd) ->
    Val = #node{type=?KV_NODE,key="fdfd",value="dfd"},
    V = base64:encode_to_string(term_to_binary(Val)),
    Len = length(V),
    ?LOG("length is ~p", [Len]),
    %?LOG("size is ~p", [Len]),
    file:pwrite(Fd, 0, V),
    file:pwrite(Fd, Len, V),
    file:sync(Fd),
    read_tree(Fd, 0).

main(Path) ->
    Val = open(?FILE_PATH),
    case Val of
        {ok, Fd} ->
            build_recursive(generatedata(10), Fd),
            ?LOG("data is ~p",[generatedata(3)]),
            ?LOG("no error");
        {error, Fd} ->
            ?LOG("Error in opening the file")
    end.
    %?LOG("~p",[Val]).

open(Path) ->
    file:open(Path, [read, raw, write, binary]).
