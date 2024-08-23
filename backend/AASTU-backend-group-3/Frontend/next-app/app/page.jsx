"use client";
import Link from "next/link";
import React from "react";
import { useState, useEffect } from "react";

const page = () => {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/blogs?limit=10")
      .then((res) => res.json())
      .then((data) => {
        setPosts(data.blogs);
        // console.log(data.blogs);
      });
  }, []);
  return (
    <div>
      <h1 className="mb-4 mt-6">
        <span className="text-2xl flex justify-center  font-semibold text-gray-800">
          Welcome to A2SV Blogger
        </span>
        <span>
          <Link href="/post/create" className="flex mr-5 my-9 justify-end">
            <p className="text-blue-500 hover:border-blue-500 hover:underline border p-2  ml-4 rounded-md">
              Create Post
            </p>
          </Link>
        </span>
      </h1>
      <ul className="list-none p-0 m-0 flex justify-center flex-wrap gap-5">
        {posts?.map((post) => (
          <li
            key={post.id}
            className="bg-white border border-gray-200 rounded-lg shadow-md p-4 mb-4 transition-shadow hover:shadow-lg w-[350px]"
          >
            <h2 className="text-2xl font-semibold text-gray-800 mb-2">
              {post.title}
            </h2>
            <p className="text-gray-700 mb-4">
              {post.content.length > 200
                ? post.content.slice(0, 200) + "..."
                : post.content}
            </p>
            <div className="flex flex-col gap-2">
              <p className="text-gray-600">
                Author: <span className="font-medium">{post.author_name}</span>
              </p>
              <p className="text-gray-600">
                <span className="font-medium text-blue-500">
                  Likes: {post.likes_count}
                </span>
                <span className="ml-4 font-medium text-red-500">
                  Dislikes: {post.dislikes_count}
                </span>
              </p>
              <p className="text-gray-600">
                Tags:{" "}
                <span className="font-medium text-gray-800">
                  {post.tags.join(", ")}
                </span>
              </p>
            </div>
            <div className="font-semibold mt-2 flex justify-end  ">
              <Link
                href={`/post/${post.id}`}
                className="text-blue-500 hover:underline"
              >
                Read More
              </Link>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default page;
