import '../../domain/entity/product.dart';

class ProductModel extends Product {
  const ProductModel({
    required String id,
    required String name,
     String? category,
    required String description,
    required String image,
    required int price,
  }) : super(
            name: name,
            category: category,
            id: id,
            description: description,
            image: image,
            price: price);

  factory ProductModel.fromJson(Map<String, dynamic> json) => ProductModel(
      category: json['category'] ?? '',
      name: json['name'],
      price: json['price'],
      id: json['id'],
      image: json['imageUrl'],
      description: json['description']);

  @override
  // ignore: override_on_non_overriding_member
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'category': category ?? '',
      'description': description,
      'price': price,
      'image': image
    };
  }

  Product toEntity() => Product(id: id, name: name, description: description, category: category, image: image, price: price);


}
