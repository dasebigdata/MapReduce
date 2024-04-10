from reducer import Reducer
from typing import List


class WordCountReducer(Reducer):

    def reduce(self, key: str, values: List[int]) -> (str, int):
        return (key, sum(values))
