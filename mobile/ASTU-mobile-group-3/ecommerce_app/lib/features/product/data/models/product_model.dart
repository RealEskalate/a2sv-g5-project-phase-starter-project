import '../../../auth/data/model/user_model.dart';
import '../../domain/entities/product.dart';

class ProductModel extends ProductEntity {
  const ProductModel(
      {required super.id,
      required super.name,
      required super.description,
      required super.price,
      required super.imageUrl,
      required super.seller});

  factory ProductModel.fromJson(Map<String, dynamic> json) {
    return ProductModel(
        id: json['id'],
        name: json['name'],
        description: json['description'],
        price: json['price'].toInt(),
        imageUrl: json['imageUrl'],
        seller: UserModel.fromSellerJson(json['seller']));
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'id': id,
      'name': name,
      'description': description,
      'price': price,
      'imageUrl': imageUrl,
      'seller': {
        '_id': seller.id,
        'name': seller.name,
        'email': seller.email,
        '__v': seller.v
      }
    };
  }

  factory ProductModel.fromEntity(ProductEntity entity) {
    return ProductModel(
        id: entity.id,
        name: entity.name,
        description: entity.description,
        price: entity.price,
        imageUrl: entity.imageUrl,
        seller: UserModel.fromEntity(entity.seller));
  }

  ProductEntity toEntity() {
    return ProductEntity(
      id: id,
      name: name,
      description: description,
      price: price,
      imageUrl: imageUrl,
      seller: UserModel.toEntityParam(seller),
    );
  }

  static List<ProductEntity> allToEntity(List<ProductModel> models) {
    List<ProductEntity> answer = <ProductEntity>[];
    for (int i = 0; i < models.length; i++) {
      answer.add(models[i].toEntity());
    }
    return answer;
  }
}
