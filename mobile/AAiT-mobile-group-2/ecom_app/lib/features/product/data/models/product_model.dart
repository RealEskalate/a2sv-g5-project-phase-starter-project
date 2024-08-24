import 'package:ecom_app/features/product/data/models/seller_model.dart';

import '../../../../core/error/exception.dart';
import '../../domain/entities/product.dart';
import '../../domain/entities/seller.dart';

class ProductModel extends Product {
  const ProductModel({
    required super.id,
    required super.name,
    required super.description,
    required super.imageUrl,
    required super.price,
    Seller? super.seller,
  });

  factory ProductModel.fromJson(Map<String, dynamic> json) {
    try {
      return ProductModel(
        id: json['id'] as String? ?? '', // Handle null case
        name: json['name'] as String? ?? '', // Handle null case
        description: json['description'] as String? ?? '', // Handle null case
        imageUrl: json['imageUrl'] as String? ?? '', // Handle null case
        price: (json['price'] as num?)?.toDouble() ?? 0.0, // Handle null case
        seller: json['seller'] != null
            ? SellerModel.fromJson(json['seller'] as Map<String, dynamic>)
            : const SellerModel(
                id: '', name: '', email: ''), // Handle null case
      );
    } catch (e) {
      throw JsonParsingException();
    }
  }

  static List<ProductModel> fromJsonList(List<dynamic> jsonList) {
    try {
      return jsonList.map((json) {
        return ProductModel.fromJson(json);
      }).toList();
    } catch (e) {
      throw JsonParsingException();
    }
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'imageUrl': imageUrl,
      'price': price,
    };
  }

  static List<Map<String, dynamic>> toJsonList(List<ProductModel> products) {
    return products.map((product) => product.toJson()).toList();
  }

  Product toEntity() => Product(
      id: id,
      name: name,
      description: description,
      imageUrl: imageUrl,
      price: price,
      seller: seller);

  static List<Product> toEntityList(List<ProductModel> models) {
    return models.map((model) => model.toEntity()).toList();
  }

  static ProductModel toModel(Product product) {
    return ProductModel(
        id: product.id,
        name: product.name,
        description: product.description,
        imageUrl: product.imageUrl,
        price: product.price,
        seller: product.seller);
  }
}
