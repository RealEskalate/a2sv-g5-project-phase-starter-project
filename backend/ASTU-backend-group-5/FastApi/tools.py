from langchain.agents import  tool

@tool
def set_decision(grade, reason) -> None:
    '''
    this tool sets the status of the blog
    it sets the grade of the blog and the reason for the decision
    this tool must always be used
    Args:
        grade : (int) (the grade of the given post from 0 to 100 )
        reason: (str) (detailed reason for the decision)
    returns:
        None
    
    '''
    from post_moderator import decision_state
    decision_state.message = reason
    decision_state.grade = grade
    return
