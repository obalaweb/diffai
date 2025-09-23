"""
PR summary generation models
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

class PRRequest(BaseModel):
    """Request for PR summary generation"""
    diff: List[Diff]
    context: Dict[str, Any]  # PR information
    options: Dict[str, Any] = {}

class PRResponse(BaseModel):
    """Response for PR summary generation"""
    success: bool
    content: str
    metadata: Dict[str, Any] = {}
    usage: Dict[str, Any] = {}
    error: Optional[str] = None
