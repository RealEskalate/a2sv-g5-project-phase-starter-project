import { useRouter } from 'next/router';

export default function ErrorPage() {
  const router = useRouter();
  const { error } = router.query;

  const errorMessages = {
    Configuration: 'There is a problem with the server configuration.',
    AccessDenied: 'You do not have permission to sign in.',
    Verification: 'The sign-in link is no longer valid. Please request a new one.',
    // Add other error types if needed
    default: 'Unable to sign in.',
  };


  return (
    <div>
      <h1>Sign in Error</h1>
      <p>{error}</p>
      {/* You can add a link back to the sign-in page or other navigation here */}
    </div>
  );
}
