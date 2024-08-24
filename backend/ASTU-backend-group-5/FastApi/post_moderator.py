from fastapi import FastAPI, HTTPException
from langchain_community.chat_models import ChatOpenAI
from langchain.schema import SystemMessage
from langchain.agents import AgentType
from langchain.agents import initialize_agent
from pydantic import BaseModel
from langchain_google_genai import ChatGoogleGenerativeAI
from tools import set_decision

import os
from dotenv import load_dotenv
load_dotenv()

llm = ChatOpenAI(temperature=0, model="gpt-4-turbo", verbose=True, api_key=os.getenv("OPENAI_API_KEY"))
print(os.getenv("OPENAI_KEY"))
class DecisionState:
    def __init__(self):
        self.grade = None
        self.message = None

decision_state = DecisionState()


system_message = SystemMessage(
    content="""You are an intelligent content moderation AI designed to analyze and validate blog posts before they are published. 
    Your primary responsibility is to ensure that every post adheres to the platform's regulations and carries a positive intent.
    You have access to tools that assist you in making accurate and informed decisions. 
    Before Your Decision You Consider the following points about the post:
    The post Shoud be
     1. as objective as possible
     2. as informative as possible
     3. as relevant as possible
     4. should agree with the title
     5. should not contain any profanity, hate speech, or insults, fake news, fake information or any other harmful content.
     and also
     As helpful AI model you have to be as polite as possible
     
     for the post given you give a point out of 100
     100 means the post is perfect and 0 means the post should never be posted and 50 means the post is not suggested to be posted you can use 
     any number from 0 to 100 to grade the post
    """
)

agent_kwargs = {
    "system_message": system_message,
}

tools = [set_decision]

moderator_agent = initialize_agent(
    tools,
    llm,
    agent=AgentType.OPENAI_FUNCTIONS,
    verbose=True,
    agent_kwargs=agent_kwargs,
    max_iterations=200,
    handle_parsing_errors=True
)
