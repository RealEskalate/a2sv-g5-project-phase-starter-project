import Image from "next/image";
import MyCards from "./components/MyCards";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24 bg-background">
      <h1>Bank dashboard</h1>
      <MyCards />
    </main>
  );
}