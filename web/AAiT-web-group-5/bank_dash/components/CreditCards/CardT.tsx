import Link from "next/link";

export default function CardT() {
  return (
    <div className="flex justify-between p-3  mb-5 mt-5">
      <p className="text-black font-extrabold">My Cards</p>
      <Link href={""} className="text-black font-extrabold">
        See All
      </Link>
    </div>
  );
}
