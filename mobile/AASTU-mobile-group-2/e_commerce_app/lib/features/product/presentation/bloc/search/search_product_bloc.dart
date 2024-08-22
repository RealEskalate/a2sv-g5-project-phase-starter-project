import 'dart:async';

import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../domain/entities/product.dart';

part 'search_product_event.dart';
part 'search_product_state.dart';

class SearchBloc extends Bloc<SearchEvent, SearchState> {
  SearchBloc() : super(SearchState()) {
    on<SearOpened>(_onOpened);
    on<ProductSearched>(_onSearched);
  }

  FutureOr<void> _onOpened(SearOpened event, Emitter<SearchState> emit) {
    final List<ProductEntity> allProducts = event.allProducts;
    emit(SearchInitial(allProducts: allProducts));
  }

  FutureOr<void> _onSearched(ProductSearched event, Emitter<SearchState> emit) {
    final List<ProductEntity> allProducts = event.allProducts;
    final name = event.name;
    final List<ProductEntity> filteredProducts =
        allProducts.where((product) => product.name.contains(name)).toList();
    emit(SearchSuccess(filtered: filteredProducts));
  }
}
