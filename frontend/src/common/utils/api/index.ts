import { AxiosError } from "axios";

const DEFAULT_ERROR_MSG = "Unable to process this action";

export function handleApiErrors(err: unknown): string {
  if (err instanceof AxiosError) {
    // Log it somewhere as well? Sentry
    return err.response ? err.response.data.message : DEFAULT_ERROR_MSG;
  }

  return err && typeof err === "object" && "message" in err ? (err as any).message : DEFAULT_ERROR_MSG;
}
