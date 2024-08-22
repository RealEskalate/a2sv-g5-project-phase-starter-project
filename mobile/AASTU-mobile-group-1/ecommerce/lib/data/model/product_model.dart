import 'package:ecommerce/core/import/import_file.dart';

// ignore: must_be_immutable
class ProductModel extends ProductEntity {
  ProductModel({
    String? id,
    required String name,
    required String imageUrl,
    required double price,
    required String description,
  }) : super(
            id: id ?? "",
            name: name,
            imageUrl: imageUrl,
            price: price,
            description: description);

  factory ProductModel.fromEntity(ProductEntity entity) {
    return ProductModel(
        id: entity.id,
        name: entity.name,
        description: entity.description,
        price: entity.price,
        imageUrl: entity.imageUrl
    );
  }

  factory ProductModel.fromjson(Map<String, dynamic> json) {
   
    return ProductModel(
        id: json['id'].toString(),
        name: json['name'].toString(),
        description: json['description'].toString(),
        price: json['price'].toDouble(),
        imageUrl: json['imageUrl'].toString());
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

  Map<String, dynamic> forApi() {
    return {
      'name': name,
      'description': description,
      'price': price,
    };
  }
}
