from typing import List, Tuple
from mapper import Mapper


class WordCountMapper(Mapper):

    def map(self, line: str) -> List[Tuple[str, int]]:
        words = line.split(" ")
        return [(word, 1) for word in words]
