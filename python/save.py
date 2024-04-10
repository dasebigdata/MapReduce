import pickle
from wordcount_mapper import WordCountMapper
from wordcount_reducer import WordCountReducer

if __name__ == '__main__':
    with open('./mapper.pkl', 'wb') as f:
        pickle.dump(WordCountMapper, file=f)
    with open('./reducer.pkl', 'wb') as f:
        pickle.dump(WordCountReducer, file=f)
