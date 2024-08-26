// components/Footer.js
import React from 'react';
import { FaTwitter, FaFacebook } from 'react-icons/fa';
import Link from 'next/link';
import { FaSquareXTwitter, FaX, FaXTwitter } from 'react-icons/fa6';

const Footer: React.FC = () => (
  <footer className="bg-gray-800 py-8 text-white">
    <div className="container mx-auto text-center">
      <div className="mb-4">
        <p>&copy; 2024 Your Bank. All rights reserved.</p>
        <p>123 Bank Street, City, Country</p>
        <p>Contact: (123) 456-7890 | Email: support@yourbank.com</p>
      </div>
      <div className="flex justify-center space-x-4">
        <Link href="https://twitter.com" passHref>
          <FaXTwitter className="h-6 w-6 text-white hover:text-blue-400" />
        </Link>
        <Link href="https://facebook.com" passHref>
          <FaFacebook className="h-6 w-6 text-white hover:text-blue-600" />
        </Link>
      </div>
    </div>
  </footer>
);

export default Footer;
