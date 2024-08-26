import 'package:e_commerce_app/features/product/data/models/product_model.dart';
import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user.dart';

class ProductEntity extends Equatable {
  String name;
  String description;
  String imageUrl;
  double price;
  String id;
  User seller;
  ProductEntity(
      {required this.description,
      required this.id,
      required this.imageUrl,
      required this.name,
      required this.price,
      required this.seller
      });
  @override
  // TODO: implement props
  List<Object?> get props => [
        name,
        id,
        description,
        imageUrl,
        price,
        seller,
      ];
  ProductModel toModel() => ProductModel(
      description: description,
      id: id,
      imageUrl: imageUrl,
      name: name,
      price: price,
      seller: User.toModel(seller)
      
      );
}
