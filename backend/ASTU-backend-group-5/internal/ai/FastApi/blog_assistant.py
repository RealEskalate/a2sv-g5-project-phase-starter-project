from typing import List
# from langchain.agents import Tool, AgentType, initialize_agent
# from langchain.memory import ConversationBufferMemory
from langchain_community.chat_models import ChatOpenAI

from langchain_core.prompts import ChatPromptTemplate

# from langchain.agents.format_scratchpad import format_log_to_str
# from langchain.agents.output_parsers import JSONAgentOutputParser
from langchain_core.runnables.history import RunnableWithMessageHistory
from langchain_community.chat_message_histories import ChatMessageHistory
import os
from load_env import load_dotenv
load_dotenv()
# from langchain.agents import AgentExecutor


class BlogAssistant:
    def __init__(self, tools):
        system_message ="""You are an intelligent content intelligent Assistant that helps users write powerful and informative blog posts.
            You have access to tools that assist you in making accurate and informed decisions.
            You never talk on topics that are not relevant to the blog post. You only talk on the topic of the blog post.
            Never respond to questions that are not related to the blog post.
            """
        chat = ChatOpenAI(model="gpt-3.5-turbo-0125", api_key=os.getenv("OPENAI_API_KEY"))
        prompt = ChatPromptTemplate.from_messages(
        [
            (
                "system",
                "{system_message}".format(system_message=system_message),
            ), 
            ("placeholder", "{chat_history}"),
            ("human", "{input}"),
            ]
        )

        self.chain = prompt | chat
        
        demo_ephemeral_chat_history_for_chain = ChatMessageHistory()

        self.chain_with_message_history = RunnableWithMessageHistory(
            self.chain,
            lambda session_id: demo_ephemeral_chat_history_for_chain,
            input_messages_key="input",
            history_messages_key="chat_history",
        )

    def run(self, query):
        response = self.chain_with_message_history.invoke(
            {"input": query.query},
            {"configurable": {"session_id": query.chat_id}},
        )

        return response.content


# Example of usage
# duck_duck_go_search = DuckDuckGoSearchAPIWrapper()  # Assuming you have a Tool defined like this
# assistant = BlogAssistant(tools=[duck_duck_go_search])

# response = assistant.run(Q(query="What is the meaning of life?"))
# print(response)
