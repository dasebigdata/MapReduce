from typing import List, Tuple, Any
from abc import abstractclassmethod, ABCMeta


class Mapper(metaclass=ABCMeta):

    @abstractclassmethod
    def map(self, line: str) -> List[Tuple[Any, Any]]:
        pass
