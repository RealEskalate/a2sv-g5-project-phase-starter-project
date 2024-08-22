part of 'search_product_bloc.dart';

class SearchEvent {
  SearchEvent();
}

class SearOpened extends SearchEvent {
  final List<ProductEntity> allProducts;
  SearOpened({required this.allProducts});
}

class ProductSearched extends SearchEvent {
  String name;
  final List<ProductEntity> allProducts;
  ProductSearched({required this.name,required this.allProducts});
}
