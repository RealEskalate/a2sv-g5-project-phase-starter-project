import 'package:equatable/equatable.dart';

import '../../../authentication/domain/entity/user.dart';
import '../../data/models/product_model.dart';

class Product extends Equatable {
  const Product({
    required this.id,
    required this.name,
    this.category,
    required this.description,
    required this.image,
    required this.price,
    required this.seller,
  });

  final String id;
  final String name;
  final String? category;
  final String description;
  final String image;
  final int price;
  final User seller;



  @override
  List<Object?> get props => [id, name, category, description, image, price, seller];

  ProductModel toModel() => ProductModel(id: id, name: name, description: description, image: image, price: price, seller: seller.toModel());

      
}
