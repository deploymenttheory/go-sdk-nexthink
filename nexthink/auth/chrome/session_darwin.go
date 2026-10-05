//go:build darwin

package chrome

import (
	"context"
	"fmt"
	"os/exec"
)

func readSession(ctx context.Context, origin string) ([]byte, error) {
	const script = `on run argv
 tell application "Google Chrome"
  repeat with w in windows
   repeat with t in tabs of w
    if URL of t starts with ((item 1 of argv) & "/") then
     return execute t javascript "(()=>{const s=JSON.parse(localStorage.getItem('okta-token-storage')||'{}').accessToken;return JSON.stringify(s?{token:s.accessToken,expiresAt:s.expiresAt}:{error:'missing'});})()"
    end if
   end repeat
  end repeat
 end tell
 return "{}"
end run`
	raw, err := exec.CommandContext(ctx, "osascript", "-e", script, origin).Output()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if err != nil {
		return nil, fmt.Errorf("cannot access Chrome session: enable macOS Automation access and Chrome JavaScript from Apple Events")
	}
	return raw, nil
}
