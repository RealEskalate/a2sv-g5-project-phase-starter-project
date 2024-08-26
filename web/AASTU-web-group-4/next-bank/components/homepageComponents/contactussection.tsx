
import React from 'react';

const ContactSection: React.FC = () => (
  <section className="py-16 text-black">
    <div className="container mx-auto text-center">
      <h2 className="text-3xl font-bold mb-4">Get in Touch</h2>
      <p className="text-lg mb-6">We would love to hear from you. Reach out to us for any queries or support.</p>
      <form className="max-w-lg mx-auto">
        <div className="mb-4">
          <input
            type="text"
            placeholder="Your Name"
            className="w-full px-4 py-2 rounded-lg border border-gray-300"
          />
        </div>
        <div className="mb-4">
          <input
            type="email"
            placeholder="Your Email"
            className="w-full px-4 py-2 rounded-lg border border-gray-300"
          />
        </div>
        <div className="mb-4">
          <textarea
            placeholder="Your Message"
            className="w-full px-4 py-2 rounded-lg border border-gray-300"
            rows={4}
          ></textarea>
        </div>
        <button
          type="submit"
          className="bg-yellow-500 text-white py-2 px-6 rounded-lg shadow-lg hover:bg-yellow-600 transition-transform transform hover:scale-105"
        >
          Send Message
        </button>
      </form>
    </div>
  </section>
);

export default ContactSection;
