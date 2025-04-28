import os
from openai import OpenAI
from dotenv import load_dotenv

load_dotenv()

def generate_insight(entry_text: str):
    # prompt = f"""
    # You are an expert project reviewer. Analyze the following entry and identify potential problems or missing parts. Respond in JSON.

    # ENTRY:
    # {entry_text}
    # """
            
    prompt = 'tell me a short story about pirates'

    client = OpenAI(
        api_key=os.environ.get("OPENAI_API_KEY"),
    )

    response = client.responses.create(
        model="gpt-4.1",
        input=prompt
    )

    print(response.output_text)

    # return response["choices"][0]["message"]["content"]

if __name__ == "__main__":
    entry = "We decided to rewrite the service layer but didn’t discuss how it affects our current tests."
    result = generate_insight(entry)
    print(result)
