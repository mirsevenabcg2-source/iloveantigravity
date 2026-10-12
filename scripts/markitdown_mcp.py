import sys
import os
from mcp.server.fastmcp import FastMCP
from markitdown import MarkItDown

mcp = FastMCP("markitdown")
md = MarkItDown()

@mcp.tool()
def convert_document(file_path: str) -> str:
    """Convert documents (PDF, DOCX, XLSX, PPTX, HTML, CSV, TXT, EPUB, etc.) to structured Markdown.
    
    Args:
        file_path: Absolute or relative path to the document file.
    """
    try:
        resolved_path = os.path.abspath(file_path)
        if not os.path.exists(resolved_path):
            return f"Error: File '{file_path}' does not exist."
        result = md.convert(resolved_path)
        return result.text_content
    except Exception as e:
        return f"Error converting document: {str(e)}"

@mcp.tool()
def convert_url(url: str) -> str:
    """Fetch and convert a web page or online document to structured Markdown.
    
    Args:
        url: URL of the webpage or online document.
    """
    try:
        result = md.convert(url)
        return result.text_content
    except Exception as e:
        return f"Error converting URL: {str(e)}"

if __name__ == "__main__":
    mcp.run()
