import Link from "next/link";

export default function CardT() {
  return (
    <div className="flex justify-between px-10 mb-5 mt-5">
      <h2 className="flex items-center text-2xl font-bold text-gray-700">
        My Cards
      </h2>
      <Link href={""} className="text-black font-extrabold">
        <p className="flex items-center text-xl font-semibold text-gray-700">
          See All
        </p>
      </Link>
    </div>
  );
}
