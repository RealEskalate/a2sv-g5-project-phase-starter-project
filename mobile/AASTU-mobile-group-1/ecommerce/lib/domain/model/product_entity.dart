import 'package:ecommerce/core/import/import_file.dart';
// ignore: must_be_immutable
class ProductEntity extends Equatable {
  String id;
  String name;
  String imageUrl;
  double price;
  String description;

  ProductEntity(
      { this.id = "",
      required this.name,
      required this.description,
      required this.price,
      required this.imageUrl});

  @override
  List<Object?> get props => [id,name,imageUrl,price,description];
}
