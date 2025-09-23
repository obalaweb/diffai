"""
Commit message generation models
"""

from pydantic import BaseModel
from typing import List, Dict, Any, Optional

class Diff(BaseModel):
    """Represents a code change"""
    path: str
    old_content: Optional[str] = None
    new_content: Optional[str] = None
    change_type: str  # added, modified, deleted, renamed
    hunk: str  # unified diff format

class CommitRequest(BaseModel):
    """Request for commit message generation"""
    diff: List[Diff]
    options: Dict[str, Any] = {}

class CommitResponse(BaseModel):
    """Response for commit message generation"""
    success: bool
    content: str
    metadata: Dict[str, Any] = {}
    usage: Dict[str, Any] = {}
    error: Optional[str] = None
