import Navbar from '@/components/homepageComponents/navbarsection';
import Hero from '@/components/homepageComponents/herosection';
import Footer from '@/components/homepageComponents/footer';
import Services from '@/components/homepageComponents/servicessection';
import FeatureSection from '@/components/homepageComponents/featuresection';
import FAQ from '@/components/homepageComponents/faqsection'; 
import Contact from '@/components/homepageComponents/contactussection';
export default function Home() {
  return (
    <>
      <Navbar/>
      <Hero />
      <Services />
      <FeatureSection />
      <FAQ/>
      <Contact/>
      <Footer />
    </>
  );
}
