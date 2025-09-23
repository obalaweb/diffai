"""
OpenAI service implementation
"""

import os
import json
import logging
from typing import List, Dict, Any, Optional
import openai
from openai import AsyncOpenAI

logger = logging.getLogger(__name__)

class OpenAIService:
    """OpenAI API service for AI operations"""
    
    def __init__(self):
        self.api_key = os.getenv("OPENAI_API_KEY")
        self.base_url = os.getenv("OPENAI_BASE_URL", "https://api.openai.com/v1")
        self.model = os.getenv("OPENAI_MODEL", "gpt-4")
        
        if self.api_key:
            self.client = AsyncOpenAI(
                api_key=self.api_key,
                base_url=self.base_url
            )
        else:
            self.client = None
            logger.warning("OpenAI API key not found")
    
    def is_available(self) -> bool:
        """Check if OpenAI service is available"""
        return self.client is not None and self.api_key is not None
    
    def get_available_models(self) -> List[str]:
        """Get list of available models"""
        return ["gpt-4", "gpt-4-turbo", "gpt-3.5-turbo"]
    
    async def generate_commit_message(
        self,
        diff: List[Dict[str, Any]],
        style: str = "conventional",
        language: str = "en",
        max_length: int = 50
    ) -> Dict[str, Any]:
        """Generate a commit message from diff"""
        if not self.is_available():
            raise Exception("OpenAI service not available")
        
        # Build prompt
        prompt = self._build_commit_prompt(diff, style, language, max_length)
        
        # Call OpenAI API
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {
                    "role": "system",
                    "content": "You are an expert software developer and Git specialist. Generate clear, concise, and professional commit messages that follow conventional commit standards."
                },
                {
                    "role": "user",
                    "content": prompt
                }
            ],
            max_tokens=200,
            temperature=0.7
        )
        
        content = response.choices[0].message.content.strip()
        
        return {
            "content": content,
            "usage": {
                "tokens": response.usage.total_tokens,
                "cost": self._calculate_cost(response.usage.total_tokens),
                "model": self.model,
                "provider": "openai"
            }
        }
    
    async def generate_pr_summary(
        self,
        pr_info: Dict[str, Any],
        diff: List[Dict[str, Any]],
        include_risk: bool = False,
        style: str = "detailed",
        language: str = "en"
    ) -> Dict[str, Any]:
        """Generate a PR summary"""
        if not self.is_available():
            raise Exception("OpenAI service not available")
        
        # Build prompt
        prompt = self._build_pr_prompt(pr_info, diff, include_risk, style, language)
        
        # Call OpenAI API
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {
                    "role": "system",
                    "content": "You are an expert software developer and code reviewer. Generate comprehensive, clear, and professional PR summaries that help developers understand changes, impact, and potential risks."
                },
                {
                    "role": "user",
                    "content": prompt
                }
            ],
            max_tokens=1000,
            temperature=0.7
        )
        
        content = response.choices[0].message.content.strip()
        
        return {
            "content": content,
            "usage": {
                "tokens": response.usage.total_tokens,
                "cost": self._calculate_cost(response.usage.total_tokens),
                "model": self.model,
                "provider": "openai"
            }
        }
    
    async def generate_changelog(
        self,
        commits: List[Dict[str, Any]],
        include_unreleased: bool = False,
        style: str = "markdown",
        language: str = "en"
    ) -> Dict[str, Any]:
        """Generate a changelog"""
        if not self.is_available():
            raise Exception("OpenAI service not available")
        
        # Build prompt
        prompt = self._build_changelog_prompt(commits, include_unreleased, style, language)
        
        # Call OpenAI API
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {
                    "role": "system",
                    "content": "You are an expert technical writer and software developer. Generate clear, well-structured changelogs that help users understand what changed, why it changed, and how it affects them."
                },
                {
                    "role": "user",
                    "content": prompt
                }
            ],
            max_tokens=2000,
            temperature=0.7
        )
        
        content = response.choices[0].message.content.strip()
        
        return {
            "content": content,
            "usage": {
                "tokens": response.usage.total_tokens,
                "cost": self._calculate_cost(response.usage.total_tokens),
                "model": self.model,
                "provider": "openai"
            }
        }
    
    def _build_commit_prompt(
        self,
        diff: List[Dict[str, Any]],
        style: str,
        language: str,
        max_length: int
    ) -> str:
        """Build prompt for commit message generation"""
        # Format diff information
        diff_summary = []
        for file_diff in diff:
            diff_summary.append(f"File: {file_diff['path']} ({file_diff['change_type']})")
            if file_diff.get('hunk'):
                diff_summary.append(f"Changes:\n{file_diff['hunk']}")
        
        diff_text = "\n\n".join(diff_summary)
        
        prompt = f"""Generate a {style} commit message for the following changes:

{diff_text}

Requirements:
- Maximum {max_length} characters
- Follow conventional commit format if style is 'conventional'
- Be clear and descriptive
- Focus on what changed and why
- Use imperative mood (e.g., "Add feature" not "Added feature")

Generate only the commit message, no additional text."""
        
        return prompt
    
    def _build_pr_prompt(
        self,
        pr_info: Dict[str, Any],
        diff: List[Dict[str, Any]],
        include_risk: bool,
        style: str,
        language: str
    ) -> str:
        """Build prompt for PR summary generation"""
        # Format PR information
        pr_text = f"""PR #{pr_info.get('number', 'N/A')}: {pr_info.get('title', 'No title')}
Author: {pr_info.get('author', 'Unknown')}
Base Branch: {pr_info.get('base_branch', 'Unknown')}
Head Branch: {pr_info.get('head_branch', 'Unknown')}"""
        
        # Format diff information
        diff_summary = []
        for file_diff in diff:
            diff_summary.append(f"File: {file_diff['path']} ({file_diff['change_type']})")
            if file_diff.get('hunk'):
                diff_summary.append(f"Changes:\n{file_diff['hunk']}")
        
        diff_text = "\n\n".join(diff_summary)
        
        prompt = f"""Generate a {style} PR summary for the following pull request:

{pr_text}

Changes:
{diff_text}

Requirements:
- Summarize what changed and why
- Highlight key features, fixes, or improvements
- Mention any breaking changes
- Include impact assessment
{"- Include risk assessment and potential issues" if include_risk else ""}
- Be clear and professional
- Use bullet points for better readability

Generate a comprehensive summary that helps reviewers understand the changes."""
        
        return prompt
    
    def _build_changelog_prompt(
        self,
        commits: List[Dict[str, Any]],
        include_unreleased: bool,
        style: str,
        language: str
    ) -> str:
        """Build prompt for changelog generation"""
        # Format commits
        commits_text = []
        for commit in commits:
            commits_text.append(f"• {commit['message']} ({commit['hash'][:8]})")
        
        commits_list = "\n".join(commits_text)
        
        prompt = f"""Generate a {style} changelog for the following commits:

{commits_list}

Requirements:
- Group changes by type (Features, Fixes, Documentation, etc.)
- Use clear, user-friendly language
- Include version information if available
{"- Include unreleased changes section" if include_unreleased else ""}
- Follow standard changelog format
- Be concise but informative
- Focus on user impact

Generate a well-structured changelog that helps users understand what changed."""
        
        return prompt
    
    def _calculate_cost(self, tokens: int) -> float:
        """Calculate approximate cost based on token usage"""
        # Rough cost calculation for GPT-4
        # This is an approximation and should be updated based on current pricing
        cost_per_1k_tokens = 0.03  # $0.03 per 1K tokens for GPT-4
        return (tokens / 1000) * cost_per_1k_tokens
