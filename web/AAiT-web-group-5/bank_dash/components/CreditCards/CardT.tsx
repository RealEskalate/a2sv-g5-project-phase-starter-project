import Link from "next/link";

export default function CardT() {
  return (
    <div className="flex justify-around px-3 mb-5 mt-5">
      <h2 className="flex items-center text-2xl font-bold text-primary-200">
        My Cards
      </h2>
      <Link href={""} className="text-black font-extrabold">
        <p className="flex items-center text-xl font-semibold text-primary-200">
          See All
        </p>
      </Link>
    </div>
  );
}
