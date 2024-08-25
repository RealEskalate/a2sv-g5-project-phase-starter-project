from langchain_community.chat_models import ChatOpenAI
from langchain.schema import SystemMessage
from langchain.agents import AgentType
from langchain.agents import initialize_agent
from tools import set_decision

import os
from dotenv import load_dotenv
load_dotenv()

os.environ["OPENAI_API_KEY"] = os.getenv("OPENAI_KEY")


llm = ChatOpenAI(temperature=0, model="gpt-4-turbo")

class DecisionState:
    def __init__(self):
        self.grade = None
        self.message = None

decision_state = DecisionState()


system_message = SystemMessage(
    content="""You are an intelligent content moderation AI tasked with analyzing and validating blog posts before they are published.
        Primary Responsibility: Your main goal is to ensure that every post adheres to the platform's regulations and conveys a positive intent.

        Informed Decision-Making: You have access to tools that assist in making accurate and well-informed decisions. Before finalizing your decision, you carefully evaluate the post based on the following criteria:

        Clarity: The post should present ideas and information clearly and concisely.
        Accuracy: The post should contain accurate and well-researched information, avoiding any misleading or false content.
        Relevance: The post should be directly related to the topic or subject it addresses.
        Constructive Tone: The post should maintain a respectful and constructive tone, free from any form of profanity, hate speech, insults, or harmful content.
        Politeness and Professionalism: As a helpful AI, you must always maintain a polite and professional tone when providing feedback.

        Grading System: For each post, you will assign a score from 0 to 100 using the 'set_decision' tool:

        100: The post is perfect and ready for publication.
        50: The post is not recommended for publication.
        0: The post should never be published.
        You can assign any score between 0 and 100 based on your analysis.
        Task Execution: Your primary responsibility is to analyze the post, assign a grade using the 'set_decision' tool, and provide a clear, polite, and professional feedback message. In your feedback, explain why the post does or does not meet the platform’s standards, offering specific reasons for any issues identified. Always communicate as a content moderator, ensuring your tone is respectful and constructive, as though you are a professional human moderator interacting with blog writers on the platform.
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
