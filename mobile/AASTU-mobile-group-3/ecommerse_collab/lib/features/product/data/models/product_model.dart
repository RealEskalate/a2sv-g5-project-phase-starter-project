import '../../../authentication/data/model/user_model.dart';
import '../../domain/entity/product.dart';

class ProductModel extends Product {
  const ProductModel({
    required String id,
    required String name,
     String? category,
    required String description,
    required String image,
    required double price,
    required UserModel seller,
  }) : super(
            name: name,
            category: category,
            id: id,
            description: description,
            image: image,
            price: price,
            seller: seller);

  factory ProductModel.fromJson(Map<String, dynamic> json) => ProductModel(
      category: json['category'] ?? '',
      name: json['name'],
      image: json['imageUrl'],
      price: (json['price']).toDouble(),
      id: json['id'],
      description: json['description'],
      seller: UserModel.fromJson(json['seller'])
      );

  @override
  // ignore: override_on_non_overriding_member
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'category': category ?? '',
      'description': description,
      'price': price,
      'image': image,
      'seller': seller,
    };
  }

  Product toEntity() => Product(id: id, name: name, description: description, category: category, image: image, price: price, seller: seller);


}
