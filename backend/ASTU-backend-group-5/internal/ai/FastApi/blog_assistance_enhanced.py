from langchain_community.chat_models import ChatOpenAI
from langchain_core.tools import tool
from langchain.agents import initialize_agent
from langchain.agents import AgentType
from langchain.schema import SystemMessage


import os
from load_env import load_dotenv
load_dotenv()

os.environ["OPENAI_API_KEY"] = os.getenv("OPENAI_API_KEY")


class Blog:
    def __init__(self, title: str, content: list):
        self.title = title
        self.content = content
     

llm = ChatOpenAI(model="gpt-3.5-turbo-0613")


@tool
def generate_image(prompt: str) -> None:
    """Useful for when you need to generate an image. for your blog post
    it generates and sets the image to the blog posts content
    args:
        prompt: (str) (what you want to generate) detailed description of the image to be generated

    """
    from openai import OpenAI
    client = OpenAI()

    response = client.images.generate(
    model="dall-e-3",
    prompt= prompt,
    size="1024x1024",
    quality="standard",
    n=1,
    )

    image_url = response.data[0].url
    blog.content.append(image_url) # adds image_url

blog = Blog("", [])
@tool
def set_title(title: str) -> None:
    """
    used to set the title of the blog post. probabliy only called once the chain starts
    
    args:
        title: (str) (the title of the blog post)
    """
    blog.title = title

@tool
def add_content(content: str) -> None:
    """
    adds content to the blog post
    the content can be a code, imae url or text
    the content must be detailed and informative
    use formatting to make it look nice
    args:
        content: (str) (the content of the blog post)
    """
    blog.content.append(content)

@tool
def add_code(code: str) -> None:
    """
    adds code to the blog post
    args:
        code: (str) (the code of the blog post)
    """
    blog.content.append(code)

class BlogAssistant:
    def __init__(self, tools):
        system_message = SystemMessage(
            content="""
                You are an advanced content creation assistant that specializes in generating dynamic and engaging blog posts similar to Medium. Your goal is to create blogs with diverse and rich content based on the queries you receive.

                Your Tasks:
                Content Creation:

                You will receive a query for a blog post and must generate a comprehensive blog entry.
                The blog post should include various types of content: text, images, code snippets, and other relevant elements.
                Blog Object:

                You will interact with a Blog object, which includes:
                Title: A string representing the blog post's title.
                Content: A list of content elements where each element can be text, images, code, or other types of media.
                Content Management:

                Diverse Content: Ensure the blog post includes different types of content (text, images, code) to provide a thorough and engaging reading experience.
                No Duplicates: Avoid adding the same content more than once.
                Rich Information: Aim to add a variety of content to ensure the blog post is informative and engaging.
                Focus:

                Only discuss topics relevant to the blog post. Do not address unrelated questions or topics.
        """
        )
        agent_kwargs = {
            "system_message": system_message,
        }
        self.tools = [tools]
        self.agent = initialize_agent(
            tools,
            llm,
            agent=AgentType.STRUCTURED_CHAT_ZERO_SHOT_REACT_DESCRIPTION,
            verbose=True,
            agent_kwargs=agent_kwargs,
            max_iterations=200,
            handle_parsing_errors=True
        )
    def run(self, query):
        self.agent.run(query)

assistant = BlogAssistant(tools=[generate_image, set_title, add_content])
assistant.run("write a blog about python. add images and sample code make it amazing")

print(blog.content)


