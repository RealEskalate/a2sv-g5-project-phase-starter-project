from langchain.agents import  tool
from langchain_community.utilities import DuckDuckGoSearchAPIWrapper
from langchain_community.chat_models import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain.agents import AgentExecutor, create_openai_tools_agent

import os
from dotenv import load_dotenv
load_dotenv()

os.environ["OPENAI_API_KEY"] = os.getenv("OPENAI_KEY")


class Blog:
    def __init__(self, title, content):
        self.title = title
        self.content = content

    def __str__(self):
        return f"Blog(title={self.title}, content={self.content})"
    
    
blog = Blog("", [])

@tool
def duck_duck_go_search(query):
    '''
    useful for when you need to answer questions about current events, trends, or advancements.
    use this tool when the knowladge you have been trained is not sufficient to answer the question
    args:
        query: (str) (what you want to search for)
    returns:
        result: (str) (the result of the search)
    '''
    search = DuckDuckGoSearchAPIWrapper()
    result = search.run(query)
    
    return result


@tool
def generate_image(prompt: str) -> str:
    """Useful for when you need to generate an image. for your blog post
    it generates and returns image uri
    the image must be detailed and informative. if the prompt is not detailed the ai will generate random image which negatively affects your blog so be careful here
    args:
        prompt: (str) (what you want to generate) detailed description of the image to be generated
    returns:
        image url: (str) (the url of the image generated)
    """
    try:
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
        return image_url
    
    except:
        return "failed to generate image"
    
    
    


@tool
def set_title(title: str) -> str:
    """
    used to set the title of the blog post. probabliy only called once the chain starts
    
    args:
        title: (str) (the title of the blog post)
    returns:
        status: (str)
    """
    blog.title = title
    return "title set success fully now the title of the blog is " + title + "keep on generating your amazing blog"

@tool
def add_content(content: str) -> str:
    """
    adds content to the blog post
    the content can be a code, imae url or text
    the content must be detailed and informative
    use markup to make it look nice
    args:
        content: (str) (the content of the blog post)
    returns:
        status: (str)
    """
    blog.content.append(content)
    return "content added succesfully"

os.environ["OPENAI_API_KEY"] = os.getenv("OPENAI_KEY")

@tool
def generate_blog_outline(blog_description: str) -> str:
    """
    generates an outline for the blog post
    you always use this tool before you generate the blog post always!
    the description must be clear and reflect the desired content of the blog
    args:
        blog_description: (str) (the description of the blog post)
    returns:
        outline: (str) (the outline of the blog post)
    """
    try:
        from openai import OpenAI
        client = OpenAI()
        response = client.chat.completions.create(
            model="gpt-4o-mini",
            messages=[
                {"role": "system", "content": """You are intellegent blog outline generator. given the blog description you generate an outline of the blog you include topics subtopics for the given description
                the outlines are generated in the following format.
                ```
                Title: <title of the blog>
                topic1: <topic1> 
                    subtopic1: <subtopic1>
                    subtopic2: <subtopic2>
                    subtopic3: <subtopic3>
                topic2: <topic2> 
                    subtopic1: <subtopic1>
                    subtopic2: <subtopic2>
                    subtopic3: <subtopic3>
                    .
                    .
                    .
                ```
                try to make the outline logical, detailed and understandable"""},
                {"role": "user", "content": "generate an outline fpr the following blog description " + blog_description},
            ]
        )
        return response.choices[0].message.content
    except:
        return "failed to generate blog outline"


class BlogAssistant:
    def __init__(self, tools):
        self.tools = [tools]
        system_message ="""You are an intelligent content intelligent Assistant that helps users write powerful and informative blog posts.
            You have access to tools that assist you in making accurate and informed decisions.
            You never talk on topics that are not relevant to the blog post. You only talk on the topic of the blog post.
            Never respond to questions that are not related to the blog post.
            You have acces to a blog object into which you can add contents by using apropirate tool.
            You always have to use generate outline tool before starting building the blog post.
            once the outline is generated you iteratively generates blog contents and adds the content to the blog object.
            try to include all contents of the generated blog outline in the blog post.
            The generated Blog must be in markup so that it can be rendered beautifully.
            Never Geberate images exccesively only generate lessthan 3 images per blog!
            If your content includes images the images should be in a markup format.
            """
        llm = ChatOpenAI(model="gpt-4-turbo", temperature=0)
            
        self.prompt = ChatPromptTemplate.from_messages(
                [
                    ("system", "{system_message}".format(system_message=system_message)),
                    MessagesPlaceholder("chat_history", optional=True),
                    ("human", "{input}"),
                    MessagesPlaceholder("agent_scratchpad"),
                ]
            )
            
        self.agent = create_openai_tools_agent(
            tools = tools,
            llm = llm,
            prompt = self.prompt,

        )
        self.agent_executor = AgentExecutor(agent=self.agent, tools=tools, verbose=True, max_iterations=30, return_intermediate_steps=True)
    def run(self, query):
        response = self.agent_executor.invoke({"input": query })
        return response

blog_assistant = BlogAssistant([generate_blog_outline, generate_image, add_content, set_title, duck_duck_go_search])