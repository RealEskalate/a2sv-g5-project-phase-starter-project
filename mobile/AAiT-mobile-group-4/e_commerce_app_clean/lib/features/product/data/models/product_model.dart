import '../../../authentication/data/model/user_model.dart';
import '../../domain/entities/product_entity.dart';

class ProductModel extends ProductEntity {
  const ProductModel({
    required super.id,
    required super.name,
    required super.description,
    required super.price,
    required super.imageUrl,
    super.seller,
  });

  factory ProductModel.fromJson(Map<String, dynamic> json) => ProductModel(
        id: json['id'],
        name: json['name'],
        description: json['description'],
        price: json['price'].toDouble(),
        imageUrl: json['imageUrl'],
        seller: UserModel.fromJson(json['seller']),
      );
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'price': price.toString(),
      'imageUrl': imageUrl,
    };
  }

  ProductEntity toProductEntity() => ProductEntity(
        id: id,
        name: name,
        description: description,
        price: price,
        imageUrl: imageUrl,
      );
  static List<ProductEntity> toProductListEntity(List<ProductModel> model) {
    return model.map((product) => product.toProductEntity()).toList();
  }
}
