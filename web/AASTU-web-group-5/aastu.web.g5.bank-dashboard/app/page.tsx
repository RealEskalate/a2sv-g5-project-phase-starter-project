// app/page.tsx
import Link from "next/link";
import Accounts from "./Accounts/page";
import Investments from "./Investments/page";

export default function Home() {
  return (
    <div>
     
      <Link href="/auth/signup">
        <button className="btn-primary bg-slate-600">Sign Up</button>
      </Link>
      <Investments />
      <Accounts />
    </div>
  );
}
