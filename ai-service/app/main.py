"""
DiffAI AI Service - FastAPI application for AI-powered Git operations
"""

from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from typing import List, Optional, Dict, Any
import logging
import os

from app.models.commit import CommitRequest, CommitResponse
from app.models.pr import PRRequest, PRResponse
from app.models.changelog import ChangelogRequest, ChangelogResponse
from services.openai_service import OpenAIService
from services.anthropic_service import AnthropicService
from services.local_service import LocalService

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Initialize FastAPI app
app = FastAPI(
    title="DiffAI AI Service",
    description="AI-powered Git assistant service",
    version="0.1.0",
    docs_url="/docs",
    redoc_url="/redoc"
)

# Add CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Initialize AI services
openai_service = OpenAIService()
anthropic_service = AnthropicService()
local_service = LocalService()

def get_ai_service(provider: str):
    """Get the appropriate AI service based on provider"""
    if provider == "openai":
        return openai_service
    elif provider == "anthropic":
        return anthropic_service
    elif provider == "local":
        return local_service
    else:
        raise HTTPException(status_code=400, detail=f"Unsupported AI provider: {provider}")

@app.get("/")
async def root():
    """Health check endpoint"""
    return {"message": "DiffAI AI Service is running", "version": "0.1.0"}

@app.get("/health")
async def health_check():
    """Detailed health check"""
    return {
        "status": "healthy",
        "version": "0.1.0",
        "services": {
            "openai": openai_service.is_available(),
            "anthropic": anthropic_service.is_available(),
            "local": local_service.is_available()
        }
    }

@app.post("/api/v1/commit", response_model=CommitResponse)
async def generate_commit_message(request: CommitRequest):
    """Generate a commit message from staged diff"""
    try:
        logger.info(f"Generating commit message for {len(request.diff)} files")
        
        # Get AI service
        ai_service = get_ai_service(request.options.get("provider", "openai"))
        
        # Generate commit message
        response = await ai_service.generate_commit_message(
            diff=request.diff,
            style=request.options.get("style", "conventional"),
            language=request.options.get("language", "en"),
            max_length=request.options.get("max_length", 50)
        )
        
        return CommitResponse(
            success=True,
            content=response["content"],
            metadata=response.get("metadata", {}),
            usage=response.get("usage", {})
        )
        
    except Exception as e:
        logger.error(f"Error generating commit message: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/api/v1/pr", response_model=PRResponse)
async def generate_pr_summary(request: PRRequest):
    """Generate a PR summary from diff and PR info"""
    try:
        logger.info(f"Generating PR summary for PR #{request.context.get('number', 'unknown')}")
        
        # Get AI service
        ai_service = get_ai_service(request.options.get("provider", "openai"))
        
        # Generate PR summary
        response = await ai_service.generate_pr_summary(
            pr_info=request.context,
            diff=request.diff,
            include_risk=request.options.get("include_risk", False),
            style=request.options.get("style", "detailed"),
            language=request.options.get("language", "en")
        )
        
        return PRResponse(
            success=True,
            content=response["content"],
            metadata=response.get("metadata", {}),
            usage=response.get("usage", {})
        )
        
    except Exception as e:
        logger.error(f"Error generating PR summary: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/api/v1/changelog", response_model=ChangelogResponse)
async def generate_changelog(request: ChangelogRequest):
    """Generate a changelog from commits"""
    try:
        logger.info(f"Generating changelog for {len(request.context)} commits")
        
        # Get AI service
        ai_service = get_ai_service(request.options.get("provider", "openai"))
        
        # Generate changelog
        response = await ai_service.generate_changelog(
            commits=request.context,
            include_unreleased=request.options.get("include_unreleased", False),
            style=request.options.get("style", "markdown"),
            language=request.options.get("language", "en")
        )
        
        return ChangelogResponse(
            success=True,
            content=response["content"],
            metadata=response.get("metadata", {}),
            usage=response.get("usage", {})
        )
        
    except Exception as e:
        logger.error(f"Error generating changelog: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@app.get("/api/v1/providers")
async def list_providers():
    """List available AI providers and their status"""
    return {
        "providers": {
            "openai": {
                "available": openai_service.is_available(),
                "models": openai_service.get_available_models()
            },
            "anthropic": {
                "available": anthropic_service.is_available(),
                "models": anthropic_service.get_available_models()
            },
            "local": {
                "available": local_service.is_available(),
                "models": local_service.get_available_models()
            }
        }
    }

if __name__ == "__main__":
    import uvicorn
    
    # Get configuration from environment
    host = os.getenv("HOST", "0.0.0.0")
    port = int(os.getenv("PORT", "8080"))
    debug = os.getenv("DEBUG", "false").lower() == "true"
    
    logger.info(f"Starting DiffAI AI Service on {host}:{port}")
    
    uvicorn.run(
        "main:app",
        host=host,
        port=port,
        reload=debug,
        log_level="info"
    )
