import '../../domain/entities/product.dart';

class ProductModel extends Product {
  const ProductModel({
    required super.id,
    required super.name,
    required super.description,
    required super.imageUrl,
    required super.price,
    required super.seller,
  });

  factory ProductModel.fromJson(Map<String, dynamic> json) {
    return ProductModel(
      id: json['id'],
      name: json['name'],
      description: json['description'],
      imageUrl: json['imageUrl'],
      price: json['price'].toDouble(),
      seller: json['seller'],
    );
  }
  @override
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'imageUrl': imageUrl,
      'price': price,
    };
  }

  Product toEntity() => Product(
      id: id,
      name: name,
      description: description,
      imageUrl: imageUrl,
      price: price, seller: seller);
  static List<Product> listToEntity(List<ProductModel> models) {
    return models.map((model) => model.toEntity()).toList();
  }
}
