from mapper import Mapper
from reducer import Reducer
import sys
import os
import pickle
from typing import List, Tuple, Any, Dict

if __name__ == '__main__':
    input_path = sys.argv[1]
    output_path = sys.argv[2]
    mapper_path = sys.argv[3]
    reducer_path = sys.argv[4]

    mapper: Mapper
    with open(mapper_path, 'rb') as f:
        mapper = pickle.load(f)()
    reducer: Reducer
    with open(reducer_path, 'rb') as f:
        reducer = pickle.load(f)()
    
    # input
    lines = []
    with open(input_path, 'r') as f:
        for line in f:
            lines.append(line[:-len(os.linesep)])
    
    # map
    pairs: List[Tuple[Any, Any]] = []
    for line in lines:
        pairs.extend(mapper.map(line))
    
    # shuffle
    intermediate: Dict[Any, list] = {}
    for pair in pairs:
        values = intermediate.get(pair[0])
        if values is None:
            intermediate[pair[0]] = [pair[1]]
        else:
            values.append(pair[1])
    intermediate_list: List[Tuple[Any, list]] = list(intermediate.items())
    intermediate_list.sort(key=lambda x: str(x[0]))

    # reduce
    result: List[Tuple[Any, Any]] = []
    for pair in intermediate_list:
        result.append(reducer.reduce(pair[0], pair[1]))

    # output
    with open(output_path, 'w') as f:
        for pair in result:
            f.write(str(pair[0]))
            f.write('\t')
            f.write(str(pair[1]))
            f.write(os.linesep)
