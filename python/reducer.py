from typing import List, Any
from abc import abstractclassmethod, ABCMeta


class Reducer(metaclass=ABCMeta):

    @abstractclassmethod
    def reduce(self, key: Any, values: list) -> (Any, Any):
        pass
