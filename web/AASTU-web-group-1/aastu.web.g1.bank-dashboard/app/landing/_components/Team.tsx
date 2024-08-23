const people = [
  {
    name: "Amanuael Kebede",
    role: "Team Lead",
    imageUrl:
      "https://storage.googleapis.com/a2sv_hub_bucket_2/images%2FAmanuael%20Kebede%20Kassie.jpg",
  },
  {
    name: "Bereket Tadiwos",
    role: "Developer",
    imageUrl:
      "https://storage.googleapis.com/a2sv_hub_bucket_2/images%2FBereket%20Tadiwos%20Dana.jpeg",
  },
  {
    name: "Emnet Teshome",
    role: "Developer",
    imageUrl:
      "https://storage.googleapis.com/a2sv_hub_bucket_2/images%2FEmnet%20Teshome%20Lulu.png",
  },
  {
    name: "Kibrnew Gedamu",
    role: "Developer",
    imageUrl:
      "https://storage.googleapis.com/a2sv_hub_bucket_2/%2Fimages%2FKibrnew%20Gedamu%20Mekonnen_G55.jpg",
  },
  {
    name: "Mohammed Shemim",
    role: "Developer",
    imageUrl:
      "https://storage.googleapis.com/a2sv_hub_bucket_2/%2Fimages%2FMohammed%20Shemim%20Awol_G55.jpg",
  },
];

export default function Team() {
  return (
    <div className="bg-white pt-24  ">
      <div className="mx-auto grid max-w-7xl gap-x-8 gap-y-20 px-6 lg:px-8 xl:grid-cols-3">
        <div className="max-w-2xl">
          <h2 className="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
            Meet our leadership
          </h2>
          <p className="mt-6 text-lg leading-8 text-gray-600">
            Our dedicated experts crafted an innovative banking dashboard,
            blending seamless functionality with intuitive design for an
            exceptional user experience.
          </p>
        </div>
        <ul
          role="list"
          className="grid gap-x-8 gap-y-12 sm:grid-cols-2 sm:gap-y-16 xl:col-span-2"
        >
          {people.map((person) => (
            <li key={person.name}>
              <div className="flex items-center gap-x-6">
                <img
                  alt=""
                  src={person.imageUrl}
                  className="h-16 w-16 rounded-full"
                />
                <div>
                  <h3 className="text-base font-semibold leading-7 tracking-tight text-gray-900">
                    {person.name}
                  </h3>
                  <p className="text-sm font-semibold leading-6 text-indigo-600">
                    {person.role}
                  </p>
                </div>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
