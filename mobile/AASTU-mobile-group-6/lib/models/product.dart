//model
import "package:flutter/cupertino.dart";
import "package:flutter/material.dart";
import "package:flutter/widgets.dart";
import 'package:provider/provider.dart';
import 'package:flutter/foundation.dart';


class Product {
  final int id;
  final double rating;
  final String imagePath;
  final String name;
  final String price;
  final String category;
  final String description;

  Product({
    required this.id,
    required this.rating,
    required this.imagePath,
    required this.name,
    required this.price,
    required this.category,
    required this.description,
  });
}




class ProductM extends ChangeNotifier {

  static int id = 4;
  static double rating = 4.5;
  List<Product> prd_list = 
  [
  Product(
    id:0,
    rating:4.1,
    imagePath: 'assets/mike.jpg',
    name: 'Travis x Nike Lows',
    description:'The Travis Scott x Nike Lows are a collaboration between the rapper Travis Scott and Nike. These low-top sneakers feature a rugged, distressed suede upper in earthy tones inspired by Travis Scott'+"'"+"s aesthetic. Design details include a reversed Swoosh, cactus-themed embellishments, and Travis Scott"+"'"+'s signature branding. The Travis Scott Lows blend high-fashion style with Nike'+"'"+"s classic low-top silhouette.",
    category: 'Men' + "'"+ 's'+' shoes',
    price: "29.99",

  ),
  Product(
    imagePath: 'assets/jeff.jpg',
    id:1,
    rating:4.1,
    name: 'Nike Mids',
    category:'Men' + "'"+ 's'+' shoes',
    description:'Nike Mids are a mid-top sneaker design from the brand. Featuring a higher ankle collar than the Lows, Mids provide additional ankle support and coverage. The mid-top silhouette retains Nike'+"'" +'s signature style, with a leather or canvas upper and cushioned midsole. Nike Mids are a popular choice for both casual wear and light athletic activities, blending comfort and support.',
    price: "39.99",
  ),Product(
      id:2,
      rating:4.1,
    imagePath: 'assets/will.jpg',
    name: 'Nike Pandas',
    description: 'Nike released a special edition "Panda" sneaker collection in 2023, featuring black and white color schemes inspired by the popular black and white panda bears. The designs incorporated panda-themed details on classic Nike silhouettes like the Air Force 1 and Dunk Low. The collection was highly sought-after by sneaker enthusiasts.',
    category:'Men ' + "'"+ 's'+' shoes',
    price: "39.99",
  ),Product(
        id:3,
rating:4.1,
    imagePath: 'assets/hightpop.jpg',
    name: 'Nike High Tops',
    category:'Men ' + "'"+ 's'+' shoes',
    description:'Nike High Tops are a classic high-top sneaker design from the brand. Featuring an extended ankle collar, High Tops provide enhanced ankle support and coverage. The high-top silhouette retains Nike'+"'"+'s signature style, with a leather or canvas upper and cushioned midsole. Nike High Tops are a timeless choice for streetwear, casual wear, and light athletic activities.',
    price: "59.99",
  ),
  // Add more items as needed
];

  get productCount => prd_list.length;
  get product => prd_list;
  List<Product> getAllData() {
    return prd_list;
  }


    
    void addProduct(Product product) {
    // product.id = id;
    id++;
    prd_list.add(product);
    notifyListeners();
  }

   void removeProduct(Product product) {
    prd_list.remove(product);
    notifyListeners();
  }
}


