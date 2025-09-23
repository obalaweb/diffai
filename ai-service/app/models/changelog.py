"""
Changelog generation models
"""

from pydantic import BaseModel
from typing import List, Dict, Any, Optional

class Commit(BaseModel):
    """Represents a commit"""
    hash: str
    author: str
    email: str
    message: str
    date: str
    files: List[str]
    insertions: int
    deletions: int

class ChangelogRequest(BaseModel):
    """Request for changelog generation"""
    context: List[Commit]  # List of commits
    options: Dict[str, Any] = {}

class ChangelogResponse(BaseModel):
    """Response for changelog generation"""
    success: bool
    content: str
    metadata: Dict[str, Any] = {}
    usage: Dict[str, Any] = {}
    error: Optional[str] = None
