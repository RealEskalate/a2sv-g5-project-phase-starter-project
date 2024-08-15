import Image from "next/image";
import Accounts from "./Accounts/page";
import Investments from "./Investments/page";

export default function Home() {
  return (
    <div>
      <Investments/>
      <Accounts />
    </div>
  );
}
