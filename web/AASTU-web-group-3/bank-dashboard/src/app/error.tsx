"use client";
import Link from "next/link";

export default function NotFound() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen px-4 py-6 bg-gray-100">
      <h1 className="text-4xl font-bold text-red-600">404</h1>
      <h2 className="mt-2 text-2xl font-semibold">Not Found</h2>
      <p className="mt-4 text-lg text-gray-700">
        Could not find the requested resource.
      </p>
      <Link href="/" className="mt-6 text-blue-500 hover:underline">
        Return Home
      </Link>
    </div>
  );
}
