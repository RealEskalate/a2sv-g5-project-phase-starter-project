part of 'search_product_bloc.dart';

class SearchState {
  
}



class SearchInitial extends SearchState {
  final List<ProductEntity> allProducts;
  SearchInitial({required this.allProducts});
}




class SearchSuccess extends SearchState{
  final List<ProductEntity> filtered;
  SearchSuccess({required this.filtered});
}
