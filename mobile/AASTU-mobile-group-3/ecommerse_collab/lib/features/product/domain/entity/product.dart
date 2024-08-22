import 'package:equatable/equatable.dart';

import '../../data/models/product_model.dart';

class Product extends Equatable {
  const Product({
    required this.id,
    required this.name,
    this.category,
    required this.description,
    required this.image,
    required this.price,
  });

  final String id;
  final String name;
  final String? category;
  final String description;
  final String image;
  final int price;

  // Map<String, dynamic> toJson() {
  //   return {
  //     'id': id,
  //     'name': name,
  //     'category': category,
  //     'description': description,
  //     'price': price,
  //     'image': image
  //   };
  // }

  @override
  List<Object?> get props => [id, name, category, description, image, price];

  ProductModel toModel() => ProductModel(id: id, name: name, description: description, image: image, price: price);

      
}
