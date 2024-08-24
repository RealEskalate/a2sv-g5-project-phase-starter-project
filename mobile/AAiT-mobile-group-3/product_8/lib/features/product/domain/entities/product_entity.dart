import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_data_entity.dart';

class Product extends Equatable {
  final String id;
  final String name;
  final String description;
  final double price;
  final String imageUrl;
  final UserDataEntity seller;
  


  const Product({
    required this.seller,
    required this.id,
    required this.name,
    required this.description,
    required this.price,
    required this.imageUrl,
   
  });

  @override
  List<Object?> get props =>
      [id, name, description, price, imageUrl, seller];
}
