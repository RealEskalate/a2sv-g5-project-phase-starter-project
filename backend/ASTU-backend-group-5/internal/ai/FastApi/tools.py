from langchain.agents import  tool
from langchain.utilities import DuckDuckGoSearchAPIWrapper

@tool
def set_decision(decision, reason):
    '''
    this tool simple returns its own argument (decision)
    if the post is valid it will be called with argument of "True"
    otherwise it is called with argument of 'False'
    so call this function with "True" or "False"
    
    Args:
        decision : (bool) (wheather the given post is valid or not)
        reason: (str) (detailed reason for the decision)
    returns:
        decision : (bool) the argument passed to the function
        reason : (str) (detailed reason for the decision)
    
    '''
    from post_moderator import decision_state
    decision_state.valid = decision
    decision_state.message = reason
    return decision


@tool
def duck_duck_go_search(query):
    '''
    useful for when you need to answer questions about current events, trends, or advancements.
    
    args:
        query: (str) (what you want to search for)
    returns:
        result: (str) (the result of the search)
    '''
    search = DuckDuckGoSearchAPIWrapper()
    result = search.run(query)
    
    return result
    