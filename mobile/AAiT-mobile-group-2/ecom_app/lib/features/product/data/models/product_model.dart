import '../../../../core/error/exception.dart';
import '../../domain/entities/product.dart';

class ProductModel extends Product {
  const ProductModel({
    required String id,
    required String name,
    required String description,
    required String imageUrl,
    required double price,
  }) : super(
            id: id,
            name: name,
            description: description,
            imageUrl: imageUrl,
            price: price);

  factory ProductModel.fromJson(Map<String, dynamic> json) {
    try {
      return ProductModel(
        id: json['id'],
        name: json['name'],
        description: json['description'],
        imageUrl: json['imageUrl'],
        price: (json['price'] as num).toDouble(),
      );
    } catch (e) {
      throw JsonParsingException();
    }
  }
  static List<ProductModel> fromJsonList(List<dynamic> jsonList) {
    try {
      return jsonList
          .map((json) => ProductModel.fromJson(json as Map<String, dynamic>))
          .toList();
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
      price: price);

  static List<Product> toEntityList(List<ProductModel> models) {
    return models.map((model) => model.toEntity()).toList();
  }

  static ProductModel toModel(Product product) {
    return ProductModel(
        id: product.id,
        name: product.name,
        description: product.description,
        imageUrl: product.imageUrl,
        price: product.price);
  }
}
