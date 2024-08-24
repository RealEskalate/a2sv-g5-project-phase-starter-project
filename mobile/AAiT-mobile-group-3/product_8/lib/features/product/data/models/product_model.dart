import '../../domain/entities/product_entity.dart';

class ProductModel extends Product {
  const ProductModel({
    required super.id,
    required super.name,
    required super.description,
    required super.price,
    required super.imageUrl,
    required super.seller,
  });

  factory ProductModel.fromJson(Map<String, dynamic> json) {
    try {
      return ProductModel(
        id: json['id'],
        name: json['name'],
        description: json['description'],
        price: (json['price'] as num).toDouble(),
        imageUrl: json['imageUrl'],
        seller: json['seller'],
      );
    } catch (e) {
      throw Exception('Error in ProductModel.fromJson: $e');
    }
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'price': price,
      'imageUrl': imageUrl,
    };
  }

  static Product toEntity(ProductModel productModel) {
    return Product(
      id: productModel.id,
      name: productModel.name,
      description: productModel.description,
      price: productModel.price,
      imageUrl: productModel.imageUrl,
      seller: productModel.seller,
    );
  }

  static List<Product> toEntityList(List<ProductModel> productModels) {
    return productModels.map((productModel) => toEntity(productModel)).toList();
  }

  static ProductModel toModel(Product product) {
    return ProductModel(
      id: product.id,
      name: product.name,
      description: product.description,
      price: product.price,
      imageUrl: product.imageUrl,
      seller: product.seller,
    );
  }

  static List<ProductModel> fromJsonList(List<dynamic> jsonList) {
    return jsonList.map((json) => ProductModel.fromJson(json)).toList();
  }

  static List<Map<String, dynamic>> toJsonList(List<ProductModel> productModels) {
    return productModels.map((productModel) => productModel.toJson()).toList();
  }
}
