#!/usr/bin/env python3
"""
ProxmoxSDK Build Script

Unified build, test, and development tool for both TypeScript and Go SDKs.
Replaces Makefile and root npm scripts with a single Python script.
"""

import argparse
import subprocess
import sys
from pathlib import Path
from typing import List, Optional


class Colors:
    """ANSI color codes for terminal output"""
    RESET = '\033[0m'
    GREEN = '\033[32m'
    RED = '\033[31m'
    YELLOW = '\033[33m'
    BLUE = '\033[34m'
    CYAN = '\033[36m'
    BOLD = '\033[1m'


def print_section(message: str, color: str = Colors.CYAN):
    """Print a formatted section header"""
    print(f"\n{color}{Colors.BOLD}{'=' * 60}{Colors.RESET}")
    print(f"{color}{Colors.BOLD}{message:^60}{Colors.RESET}")
    print(f"{color}{Colors.BOLD}{'=' * 60}{Colors.RESET}\n")


def print_success(message: str):
    """Print a success message"""
    print(f"{Colors.GREEN}✓ {message}{Colors.RESET}")


def print_error(message: str):
    """Print an error message"""
    print(f"{Colors.RED}✗ {message}{Colors.RESET}", file=sys.stderr)


def print_info(message: str):
    """Print an info message"""
    print(f"{Colors.BLUE}→ {message}{Colors.RESET}")


def run_command(cmd: List[str], cwd: Optional[Path] = None, check: bool = True) -> bool:
    """Run a shell command and return success status"""
    try:
        print_info(f"Running: {' '.join(cmd)}")
        result = subprocess.run(
            cmd,
            cwd=cwd,
            check=check,
            capture_output=False
        )
        return result.returncode == 0
    except subprocess.CalledProcessError as e:
        print_error(f"Command failed with exit code {e.returncode}")
        return False
    except FileNotFoundError:
        print_error(f"Command not found: {cmd[0]}")
        return False


class SDKBuilder:
    """Main builder class for ProxmoxSDK monorepo"""
    
    def __init__(self):
        self.root = Path(__file__).parent
        self.ts_dir = self.root / "typescript"
        self.go_dir = self.root / "go"
    
    def install(self) -> bool:
        """Install dependencies for both SDKs"""
        print_section("Installing Dependencies")
        
        success = True
        
        # TypeScript dependencies
        print_info("Installing TypeScript dependencies...")
        if self.ts_dir.exists():
            if run_command(["npm", "install"], cwd=self.ts_dir):
                print_success("TypeScript dependencies installed")
            else:
                print_error("Failed to install TypeScript dependencies")
                success = False
        
        # Go dependencies
        print_info("Installing Go dependencies...")
        if self.go_dir.exists():
            if run_command(["go", "mod", "download"], cwd=self.go_dir):
                print_success("Go dependencies installed")
            else:
                print_error("Failed to install Go dependencies")
                success = False
        
        return success
    
    def build_ts(self) -> bool:
        """Build TypeScript SDK"""
        print_section("Building TypeScript SDK")
        if not self.ts_dir.exists():
            print_error("TypeScript directory not found")
            return False
        
        if run_command(["npm", "run", "build"], cwd=self.ts_dir):
            print_success("TypeScript SDK built successfully")
            return True
        else:
            print_error("TypeScript build failed")
            return False
    
    def build_go(self) -> bool:
        """Build Go SDK"""
        print_section("Building Go SDK")
        if not self.go_dir.exists():
            print_error("Go directory not found")
            return False
        
        if run_command(["go", "build", "./..."], cwd=self.go_dir):
            print_success("Go SDK built successfully")
            return True
        else:
            print_error("Go build failed")
            return False
    
    def build(self) -> bool:
        """Build both SDKs"""
        ts_ok = self.build_ts()
        go_ok = self.build_go()
        
        if ts_ok and go_ok:
            print_section("Build Complete", Colors.GREEN)
            return True
        else:
            print_section("Build Failed", Colors.RED)
            return False
    
    def test_ts(self) -> bool:
        """Run TypeScript tests"""
        print_section("Running TypeScript Tests")
        if not self.ts_dir.exists():
            print_error("TypeScript directory not found")
            return False
        
        if run_command(["npm", "test", "run"], cwd=self.ts_dir):
            print_success("TypeScript tests passed")
            return True
        else:
            print_error("TypeScript tests failed")
            return False
    
    def test_go(self) -> bool:
        """Run Go tests"""
        print_section("Running Go Tests")
        if not self.go_dir.exists():
            print_error("Go directory not found")
            return False
        
        if run_command(["go", "test", "-v", "./..."], cwd=self.go_dir):
            print_success("Go tests passed")
            return True
        else:
            print_error("Go tests failed")
            return False
    
    def test(self) -> bool:
        """Run all tests"""
        ts_ok = self.test_ts()
        go_ok = self.test_go()
        
        if ts_ok and go_ok:
            print_section("All Tests Passed", Colors.GREEN)
            return True
        else:
            print_section("Some Tests Failed", Colors.RED)
            return False
    
    def lint_ts(self) -> bool:
        """Lint TypeScript code"""
        print_section("Linting TypeScript")
        if not self.ts_dir.exists():
            print_error("TypeScript directory not found")
            return False
        
        if run_command(["npx", "@biomejs/biome", "check", "."], cwd=self.ts_dir):
            print_success("TypeScript linting passed")
            return True
        else:
            print_error("TypeScript linting failed")
            return False
    
    def lint_go(self) -> bool:
        """Lint Go code"""
        print_section("Linting Go")
        if not self.go_dir.exists():
            print_error("Go directory not found")
            return False
        
        if run_command(["go", "vet", "./..."], cwd=self.go_dir):
            print_success("Go linting passed")
            return True
        else:
            print_error("Go linting failed")
            return False
    
    def lint(self) -> bool:
        """Lint all code"""
        ts_ok = self.lint_ts()
        go_ok = self.lint_go()
        return ts_ok and go_ok
    
    def format_ts(self) -> bool:
        """Format TypeScript code"""
        print_section("Formatting TypeScript")
        if not self.ts_dir.exists():
            print_error("TypeScript directory not found")
            return False
        
        if run_command(["npx", "@biomejs/biome", "format", "--write", "."], cwd=self.ts_dir):
            print_success("TypeScript code formatted")
            return True
        else:
            print_error("TypeScript formatting failed")
            return False
    
    def format_go(self) -> bool:
        """Format Go code"""
        print_section("Formatting Go")
        if not self.go_dir.exists():
            print_error("Go directory not found")
            return False
        
        if run_command(["go", "fmt", "./..."], cwd=self.go_dir):
            print_success("Go code formatted")
            return True
        else:
            print_error("Go formatting failed")
            return False
    
    def format(self) -> bool:
        """Format all code"""
        ts_ok = self.format_ts()
        go_ok = self.format_go()
        return ts_ok and go_ok
    
    def clean_ts(self) -> bool:
        """Clean TypeScript build artifacts"""
        print_section("Cleaning TypeScript")
        if not self.ts_dir.exists():
            print_error("TypeScript directory not found")
            return False
        
        if run_command(["npm", "run", "clean"], cwd=self.ts_dir):
            print_success("TypeScript cleaned")
            return True
        else:
            print_error("TypeScript clean failed")
            return False
    
    def clean_go(self) -> bool:
        """Clean Go build artifacts"""
        print_section("Cleaning Go")
        if not self.go_dir.exists():
            print_error("Go directory not found")
            return False
        
        if run_command(["go", "clean"], cwd=self.go_dir):
            print_success("Go cleaned")
            return True
        else:
            print_error("Go clean failed")
            return False
    
    def clean(self) -> bool:
        """Clean all build artifacts"""
        ts_ok = self.clean_ts()
        go_ok = self.clean_go()
        return ts_ok and go_ok


def main():
    """Main entry point"""
    parser = argparse.ArgumentParser(
        description="ProxmoxSDK Build Tool",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  ./build.py install           Install all dependencies
  ./build.py build             Build both SDKs
  ./build.py build --ts        Build TypeScript SDK only
  ./build.py test              Run all tests
  ./build.py test --go         Run Go tests only
  ./build.py lint              Lint all code
  ./build.py format            Format all code
  ./build.py clean             Clean all build artifacts
        """
    )
    
    parser.add_argument(
        'command',
        choices=['install', 'build', 'test', 'lint', 'format', 'clean'],
        help='Command to execute'
    )
    
    parser.add_argument(
        '--ts', '--typescript',
        action='store_true',
        help='Target TypeScript SDK only'
    )
    
    parser.add_argument(
        '--go',
        action='store_true',
        help='Target Go SDK only'
    )
    
    args = parser.parse_args()
    
    builder = SDKBuilder()
    
    # Determine which SDK to target
    target_both = not (args.ts or args.go)
    
    # Execute command
    success = True
    
    if args.command == 'install':
        success = builder.install()
    
    elif args.command == 'build':
        if args.ts:
            success = builder.build_ts()
        elif args.go:
            success = builder.build_go()
        else:
            success = builder.build()
    
    elif args.command == 'test':
        if args.ts:
            success = builder.test_ts()
        elif args.go:
            success = builder.test_go()
        else:
            success = builder.test()
    
    elif args.command == 'lint':
        if args.ts:
            success = builder.lint_ts()
        elif args.go:
            success = builder.lint_go()
        else:
            success = builder.lint()
    
    elif args.command == 'format':
        if args.ts:
            success = builder.format_ts()
        elif args.go:
            success = builder.format_go()
        else:
            success = builder.format()
    
    elif args.command == 'clean':
        if args.ts:
            success = builder.clean_ts()
        elif args.go:
            success = builder.clean_go()
        else:
            success = builder.clean()
    
    sys.exit(0 if success else 1)


if __name__ == '__main__':
    main()
