package users

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/zwoooz/gator/internal/cli"
)

func reset(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("%w: expected 0 argument, found %v", cli.ErrUsage, len(cmd.Args))
	}

	ok, err := confirm("This will delete ALL users. Continue?")
	if err != nil {
		return err
	}
	if !ok {
		fmt.Println("Aborted.")
		return nil
	}

	err = s.DB.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Delted all rows in users table")

	return nil
}

func confirm(prompt string) (bool, error) {
	fmt.Printf("%s [y/N]: ", prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return false, scanner.Err()
	}

	answer := strings.ToLower(scanner.Text())
	return answer == "y" || answer == "yes", nil
}
