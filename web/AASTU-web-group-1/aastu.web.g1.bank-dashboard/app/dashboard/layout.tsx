import Header from "./_components/Header";
import Sidebar from "./_components/Sidebar";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <main className="flex gap-5">
      <Sidebar />
      <div>
        <Header />
        {children}
      </div>
    </main>
  );
}
