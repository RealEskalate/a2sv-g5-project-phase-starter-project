import '../../domain/entitity/product.dart';
import '../../domain/entitity/user.dart';
import 'user_model.dart';

class ProductModel extends Product {
  const ProductModel({
    required String id,
    required String name,
    required double price,
    required String description,
    required String imageUrl,
    required User seller,
  }) : super(
          id: id,
          name: name,
          price: price,
          description: description,
          imageUrl: imageUrl,
          seller: seller,
        );

  factory ProductModel.fromJson(Map<String, dynamic> json) {
    return ProductModel(
      id: json['id'],
      name: json['name'],
      price: json['price'].toDouble(),
      description: json['description'],
      imageUrl: json['imageUrl'],
      seller: UserModel.fromJson(json['seller']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'price': price,
      'description': description,
      'imageUrl': imageUrl,
      'seller':seller,
    };
  }

  factory ProductModel.fromProduct(Product product) {
    return ProductModel(
      id: product.id,
      name: product.name,
      price: product.price,
      description: product.description,
      imageUrl: product.imageUrl,
      seller: product.seller,
    );
  }
}
