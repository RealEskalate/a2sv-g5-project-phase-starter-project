import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import '../../../domain/entities/product.dart';
import '../../../domain/use_case/get_all_products.dart';

part 'search_page_event.dart';
part 'search_page_state.dart';

class SearchPageBloc extends Bloc<SearchPageEvent, SearchPageState> {
  final GetAllProducts getAllProducts;
   List<Product> orginal = [];

  SearchPageBloc({required this.getAllProducts}) : super(SearchPageInitialState()) {
    on<FetchAllProductsSearchEvent>(_onFetchSearchAllProducts);
    on<SearchProductsEvent>(_onSearchProducts);
  }

  Future<void> _onFetchSearchAllProducts(FetchAllProductsSearchEvent event, Emitter<SearchPageState> emit) async {
    emit(SearchPageLoadingState());

    final result = await getAllProducts();
    result.fold(
      (failure) => emit(const SearchPageErrorState('Failed to fetch products')),
      (products)  {
        orginal = products;
        emit(SearchPageLoadedState(products));
        },
    );
  }

  Future<void> _onSearchProducts(SearchProductsEvent event, Emitter<SearchPageState> emit) async {
    emit(SearchPageLoadingState());

    final result = orginal;
    if(event.query.trim() == '') {
     emit(SearchPageLoadedState(orginal));
    }
    
        final filteredProducts = result.where((product) =>
          product.name.trim().toLowerCase().contains(event.query.trim().toLowerCase())).toList();
        emit(SearchPageLoadedState(filteredProducts));
    
  }
}
